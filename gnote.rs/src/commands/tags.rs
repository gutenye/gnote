use clap::Args;
use std::fs;
use std::path::Path;
use std::process;

/// Generate tags file
#[derive(Args, Debug)]
pub struct Tags {
	/// Input note directory
	#[clap(long, default_value_t = String::from("~/env/note"))]
	dir: String,

	/// Ouput tags file
	#[clap(long, default_value_t = String::from("~/tags"))]
	output: String,

	/// Marker string
	#[clap(long, default_value_t = String::from("∗"))]
	marker: String,

	/// Cache directory
	#[clap(long, default_value_t = String::from("~/.cache/gnote"))]
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
		self.create_tags();
	}

	fn create_tags(&self) {
		self.check_dirs();
	}

	fn check_dirs(&self) {
		if !Path::new(&self.dir).exists() {
			eprintln!("Note directory not found: '{}'", self.dir);
			process::exit(1);
		}
	}
}
