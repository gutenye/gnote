mod commands;

use clap::{Parser, Subcommand};

#[derive(Parser, Debug)]
#[clap(about, version, author)]
struct Cli {
  #[clap(subcommand)]
  command: Commands,
}

#[derive(Subcommand, Debug)]
enum Commands {
  Tags(commands::tags::Tags),
}

fn main() {
  let cli = Cli::parse();
  match &cli.command {
    Commands::Tags(tags) => tags.execute(),
  }
}
