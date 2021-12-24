import fs from 'fs/promises'
import glob from 'fast-glob'
import { escapeRegExp } from 'lodash-es'
import { exit, writeFileWithMkdir } from '@/utils'

export interface TagsContext {
  dir: string
  output: string
  marker: string
  cache: string
  watch: boolean
  extension: string
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

/**
 * read a.gnote and return tags format
 */
export async function extractTagsFromFile(
  file: string,
  { dir, marker }: TagsContext
): Promise<string> {
  const fullFile = `${dir}/${file}`
  const text = await fs.readFile(fullFile, 'utf8')
  return extractTagsFromText({ text, path: fullFile, marker })
}

/**
 * List **\/*.gnote
 */
export async function listFiles({
  dir,
  extension,
}: TagsContext): Promise<string[]> {
  return await glob(`**/*${extension}`, { cwd: dir })
}

/**
 * Create ~/.cache/gnote/a.gnote
 */
export async function writeTagsToCache(
  file: string,
  tags: string,
  { cache }: TagsContext
): Promise<void> {
  if (tags === '') {
    return
  }
  await writeFileWithMkdir(`${cache}/${file}`, tags)
}

/**
 * create ~/tags
 */
export async function createAllTagsFromCache({
  cache,
  extension,
  output,
}: TagsContext): Promise<void> {
  const files = await glob(`${cache}/**/*${extension}`)
  let allTags = ''
  for (const file of files) {
    const tags = await fs.readFile(file, 'utf8')
    allTags += `\n${tags}`
  }
  allTags = sortTags(allTags)
  const result = `!_TAG_FILE_SORTED	1\n${allTags}`
  await writeFileWithMkdir(output, result)
}

function sortTags(text: string): string {
  return text
    .split(/\n/)
    .sort((a, b) => a.localeCompare(b))
    .join('\n')
}

/**
 * Empty ~/.cache/gnote directory
 */
export async function emptyCacheDir({ cache }: TagsContext): Promise<void> {
  const files = await fs.readdir(cache)
  for (const file of files) {
    await fs.rm(`${cache}/${file}`, { recursive: true })
  }
}

export async function checksDir({ dir, cache }: TagsContext): Promise<void> {
  try {
    await fs.access(dir)
  } catch (err) {
    if (err.code === 'ENOENT') {
      exit(`Note directory not found: '${dir}'`)
    }
    return err
  }
}
