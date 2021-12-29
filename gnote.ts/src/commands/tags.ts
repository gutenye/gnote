import fs from 'fs/promises'
import glob from 'fast-glob'
import chokidar from 'chokidar'
import { escapeRegExp } from 'lodash-es'
import { Command, Option } from 'clipanion'
import { exit, writeFileWithMkdir, emptyDir } from '#/utils'

const { HOME } = process.env as { [key: string]: string }

export default class Tags extends Command {
  static paths = [['tags']]

  static usage = Command.Usage({
    description: 'Generate tags file',
    details: `
      - \`--note-dir\` (~/env/note)
      - \`--note-marker\` (∗)
      - \`--note-extension\` (.gnote)
      - \`--output\` (~/tags)
      - \`--cache-dir\` (~/.cache/gnote)
      - \`--watch\` (false)
      
      It will empty the cache directory first.
    `,
    examples: [
      ['Generate tags', 'gnote tags'],
      ['Generate tags and watch', 'gnote tags --watch'],
    ],
  })

  noteDir = Option.String('--note-dir', `${HOME}/env/note`, {
    description: 'Input note directory',
  })

  noteMarker = Option.String('--note-marker', '∗', {
    description: 'Marker string',
  })

  noteExtension = Option.String('--note-extension', '.gnote', {
    description: 'Note extension',
  })

  output = Option.String('--output', `${HOME}/tags`, {
    description: 'Ouput tags file',
  })

  cacheDir = Option.String('--cache-dir', `${HOME}/.cache/gnote`, {
    description: 'Cache directory',
  })

  watch = Option.Boolean('--watch', false, {
    description: 'Watch mode',
  })

  async execute(): Promise<void> {
    await this.createTags()

    if (this.watch) {
      this.startWatch()
    }
  }

  async createTags(): Promise<void> {
    await this.checksDir()

    await this.emptyCacheDir()

    const notes = await this.listNotes()
    for (const note of notes) {
      await this.createTagsInCache(note)
    }

    await this.createAllTagsFromCache()

    console.log(`Created ${this.output}`)
  }

  async checksDir(): Promise<void> {
    try {
      await fs.access(this.noteDir)
    } catch (err) {
      if (err.code === 'ENOENT') {
        exit(`Note directory not found: '${this.noteDir}'`)
      }
      return err
    }
  }

  /**
   * Empty ~/.cache/gnote directory
   */
  async emptyCacheDir(): Promise<void> {
    await fs.mkdir(this.cacheDir, { recursive: true })
    await emptyDir(this.cacheDir)
  }

  /**
   * List **\/*.gnote
   */
  async listNotes(): Promise<string[]> {
    return await glob(`**/*${this.noteExtension}`, { cwd: this.noteDir })
  }

  /**
   * Create <cacheDir>/a.gnote
   */
  async createTagsInCache(note: string): Promise<void> {
    const tags = await this.extractTagsFromFile(note)
    await this.writeTagsToCache(note, tags)
  }

  /**
   * read <noteDir>/a.gnote and return tags content
   */
  async extractTagsFromFile(note: string): Promise<string> {
    const fullFile = `${this.noteDir}/${note}`
    const text = await fs.readFile(fullFile, 'utf8')
    return extractTagsFromText({
      text,
      path: fullFile,
      marker: this.noteMarker,
    })
  }

  /**
   * Create ~/.cache/gnote/a.gnote
   */
  async writeTagsToCache(note: string, tags: string): Promise<void> {
    if (tags === '') {
      return
    }
    await writeFileWithMkdir(`${this.cacheDir}/${note}`, tags)
  }

  /**
   * create ~/tags
   */
  async createAllTagsFromCache(): Promise<void> {
    const files = await glob(`${this.cacheDir}/**/*${this.noteExtension}`)
    let allTags = ''
    for (const file of files) {
      const tags = await fs.readFile(file, 'utf8')
      allTags += `\n${tags}`
    }
    allTags = sortTags(allTags)
    const result = `!_TAG_FILE_SORTED	1\n${allTags}`
    await writeFileWithMkdir(this.output, result)
  }

  startWatch(): void {
    const watcher = chokidar.watch(`**/*${this.noteExtension}`, {
      cwd: this.noteDir,
    })
    watcher.on('change', async function (file: string) {
      console.log(`Changed ${file}`)
      await this.createTagsInCache(file)
      await this.createAllTagsFromCache()
    })
    console.log('Watching for changes...')
  }
}

/**
 * Read text and returns tags format
 *
 * Tags Format
 *   hello	/path/hello.gnote /∗hello∗
 */
export function extractTagsFromText({
  text,
  path,
  marker,
}: {
  text: string
  path: string
  marker: string
}): string {
  const pattern = new RegExp(
    `${escapeRegExp(marker)}([^\n ]+)${escapeRegExp(marker)}`,
    'g'
  )
  const matches = text.matchAll(pattern)
  const ids = Array.from(matches).map((v) => v[1])
  const tags = ids
    .map((id) => {
      const jump = `/${marker}${id}${marker}`
      return `${id}\t${path}\t${jump}`
    })
    .join('\n')
  return tags
}

function sortTags(text: string): string {
  return text
    .split(/\n/)
    .sort((a, b) => a.localeCompare(b))
    .join('\n')
}
