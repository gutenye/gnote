use crate::utils;
use clap::Args;
use glob::glob;
use regex::Regex;
use shellexpand::tilde;
use std::fs;
use std::path::{Path, PathBuf};
use std::process;

/// Generate tags file
#[derive(Args, Debug)]
pub struct TagsContext {
	/// Note directory
	#[clap(long="note-dir", default_value_t=String::from(tilde("~/env/note")))]
	note_dir: String,

	/// Note extension
	#[clap(long="note-extension", default_value_t=String::from(".gnote"))]
	note_extension: String,

	/// Note marker string
	#[clap(long="note-marker", default_value_t=String::from("∗"))]
	note_marker: String,

	/// Ouput tags file
	#[clap(long, default_value_t=String::from(tilde("~/tags")))]
	output: String,

	/// Cache directory
	#[clap(long="cache-dir", default_value_t=String::from(tilde("~/.cache/gnote")))]
	cache_dir: String,

	/// Watch mode
	#[clap(long)]
	watch: bool,
}

pub struct Tags {
	note_dir: PathBuf,
	note_extension: String,
	note_marker: String,
	output: PathBuf,
	cache_dir: PathBuf,
	watch: bool,
	pattern: Regex,
}

impl Tags {
	pub fn new(context: &TagsContext) -> Self {
		Self {
			note_dir: PathBuf::from(&context.note_dir),
			note_extension: context.note_extension.to_string(),
			note_marker: context.note_marker.to_string(),
			output: PathBuf::from(&context.output),
			cache_dir: PathBuf::from(&context.cache_dir),
			watch: context.watch,
			pattern: create_pattern(&context.note_marker),
		}
	}

	pub fn execute(&self) {
		self.create_tags();
	}

	/// Create tags file
	fn create_tags(&self) {
		self.check_dirs();

		self.empty_cache_dir();

		let note_paths = self.list_notes();
		for note_path in note_paths {
			self.create_tags_in_cache(&note_path)
		}

		self.create_all_tags_from_cache();

		println!("Created {}", self.output.display());
	}

	/// Make sure note directory exists
	fn check_dirs(&self) {
		if !self.note_dir.exists() {
			eprintln!("Note directory not found: '{}'", self.note_dir.display());
			process::exit(1);
		}
	}

	/// Empty cache directory
	fn empty_cache_dir(&self) {
		fs::create_dir_all(&self.cache_dir).expect(&format!(
			"Failed to create cache dir: {}",
			self.cache_dir.display()
		));
		utils::empty_dir(&self.cache_dir).expect(&format!(
			"Failed to empty cache dir: {}",
			self.cache_dir.display()
		));
	}

	fn list_notes(&self) -> Vec<PathBuf> {
		glob(&format!(
			"{}/**/*{}",
			self.note_dir.to_string_lossy(),
			self.note_extension
		))
		.unwrap()
		.map(|file| {
			file
				.unwrap()
				.strip_prefix(&self.note_dir)
				.unwrap()
				.to_path_buf()
		})
		.collect()
	}

	/// Create <cache>/a.gnote
	fn create_tags_in_cache(&self, note_path: &Path) {
		let tags_content = self.extract_tags_from_file(note_path);
		self.write_tags_to_cache(note_path, &tags_content);
	}

	/// Read <note>/a.gnote and returns tags content
	fn extract_tags_from_file(&self, note_path: &Path) -> String {
		let full_node_path = &self.note_dir.join(note_path);
		let content = fs::read_to_string(&full_node_path).expect(&format!(
			"Failed to read file: {}",
			full_node_path.display()
		));
		extract_tags_from_text(&content, &full_node_path, &self.note_marker, &self.pattern)
	}

	fn write_tags_to_cache(&self, note_path: &Path, tags_content: &str) {
		if tags_content.is_empty() {
			return;
		}
		let full_note_cache_path = self.cache_dir.join(note_path);
		utils::write_with_create_dir(&full_note_cache_path, tags_content).expect(&format!(
			"Failed to write tags to cache: {}",
			note_path.display()
		));
	}

	fn create_all_tags_from_cache(&self) {
		let paths = glob(&format!(
			"{}/**/*{}",
			self.cache_dir.to_string_lossy(),
			self.note_extension
		))
		.unwrap();
		let mut all_tags_content = String::new();
		for path in paths {
			let path = path.unwrap();
			let tags_content = fs::read_to_string(&path).unwrap();
			all_tags_content.push_str(&format!("\n{}", &tags_content));
		}
		let all_tags_content = sort_tags(&all_tags_content);
		let result = format!("!_TAG_FILE_SORTED\t1\n{}", &all_tags_content);
		utils::write_with_create_dir(&self.output, &result).expect(&format!(
			"Failed to write tags to file: {}",
			self.output.display()
		));
	}
}

fn create_pattern(note_marker: &str) -> Regex {
	Regex::new(&format!(r"{0}([^\s]+){0}", regex::escape(&note_marker))).unwrap()
}

fn extract_tags_from_text(text: &str, path: &Path, note_marker: &str, pattern: &Regex) -> String {
	let ids: Vec<&str> = pattern
		.captures_iter(&text)
		.map(|cap| cap.get(1).unwrap().as_str())
		.collect();

	let tags_content = ids
		.iter()
		.map(|id| {
			let jump = format!("/{0}{1}{0}", note_marker, id);
			format!("{}\t{}\t{}", id, path.to_string_lossy(), jump)
		})
		.collect::<Vec<String>>()
		.join("\n");

	return tags_content;
}

#[cfg(test)]
mod tests {
	use super::*;

	#[test]
	fn test_extract_tags_from_text() {
		let result = extract_tags_from_text(
			"
				hello world
				foo *bar*
				link step
				*baz* and 
			",
			Path::new("/a.gnote"),
			"*",
			&create_pattern("*"),
		);
		assert_eq!(
			result,
			"\
				bar\t/a.gnote\t/*bar*\n\
				baz\t/a.gnote\t/*baz*\
			"
		)
	}
}

fn sort_tags(tags_content: &str) -> String {
	let mut lines: Vec<_> = tags_content.lines().collect();
	lines.sort();
	lines.join("\n")
}
