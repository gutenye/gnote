use clap::Args;
use std::fs;

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
		println!("{}", self.dir);
		// fs::try_exists(&self.dir).expect("Error");
		fs::try_exists("/tmp/a/adsfoadf/afdaf").expect("Error");
		// .expect(format!("Note directory not found: '{dir}'", dir = self.dir).as_str());
	}
}
