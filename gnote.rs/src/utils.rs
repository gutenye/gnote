use std::fs;
use std::io;
use std::path::Path;

pub fn empty_dir<P: AsRef<Path>>(path: P) -> io::Result<()> {
	for entry in fs::read_dir(path)? {
		let entry = entry?;
		let path = entry.path();
		if path.is_dir() {
			fs::remove_dir_all(path)?
		} else {
			fs::remove_file(path)?
		}
	}
	Ok(())
}
