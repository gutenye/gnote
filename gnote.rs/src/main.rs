mod commands;
mod utils;

use clap::{Parser, Subcommand};
use commands::tags::{Tags, TagsContext};

#[derive(Parser, Debug)]
#[clap(about, version, author)]
struct Cli {
  #[clap(subcommand)]
  command: Commands,
}

#[derive(Subcommand, Debug)]
enum Commands {
  Tags(TagsContext),
}

fn main() {
  let cli = Cli::parse();
  match &cli.command {
    Commands::Tags(tags_context) => Tags::new(tags_context).execute(),
  }
}
