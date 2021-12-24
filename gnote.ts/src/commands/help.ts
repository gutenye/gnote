import { Command, Option } from 'clipanion'

export default class Help extends Command {
  static paths = [['help']]

  rest = Option.Proxy()

  async execute(): Promise<void> {
    await this.cli.run([...this.rest, '--help'])
  }
}
