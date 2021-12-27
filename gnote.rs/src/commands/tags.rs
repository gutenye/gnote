use crate::utils;
use clap::Args;
use shellexpand::tilde;
use std::fs;
use std::path::Path;
use std::process;

/// Generate tags file
#[derive(Args, Debug)]
pub struct Tags {
	/// Input note directory
	#[clap(long, default_value_t = String::from(tilde("~/env/note")))]
	dir: String,

	/// Ouput tags file
	#[clap(long, default_value_t = String::from(tilde("~/tags")))]
	output: String,

	/// Marker string
	#[clap(long, default_value_t = String::from("∗"))]
	marker: String,

	/// Cache directory
	#[clap(long, default_value_t = String::from(tilde("~/.cache/gnote")))]
	cache: String,

	/// Note extension
	#[clap(long, default_value_t = String::from(".gnote"))]
	extension: String,

	/// Watch mode
	#[clap(long)]
	watch: bool,
}

impl Tags {
	pub fn execute(&self) {
		// println!("{:?}", self);
		self.create_tags();
	}

	fn create_tags(&self) {
		self.check_dirs();
		self.empty_cache_dir();
	}

	fn check_dirs(&self) {
		if !Path::new(&self.dir).exists() {
			eprintln!("Note directory not found: '{}'", self.dir);
			process::exit(1);
		}
	}

	fn empty_cache_dir(&self) {
		fs::create_dir_all(&self.cache).expect(&format!("Failed to create cache dir '{}'", self.cache));
		utils::empty_dir(&self.cache).expect(&format!("Failed to empty cache dir '{}'", self.cache));
	}
}
