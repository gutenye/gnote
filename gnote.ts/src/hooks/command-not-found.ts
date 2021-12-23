import { Hook, Help } from '@oclif/core'

const notFoundHook: Hook<'not-found'> = async function (options) {
  new Help(this.config).showHelp([])
}

export default notFoundHook