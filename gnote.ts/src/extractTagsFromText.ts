import fs from 'fs-jetpack'
import { escapeRegExp } from 'lodash'

// hello	/path/hello.gnote /∗hello∗

interface Options {
  text: string
  path: string
  marker: string
}

export default function extractTagsFromFile({
  text,
  path,
  marker,
}: Options): string {
  const pattern = new RegExp(
    `${escapeRegExp(marker)}([^\n ]+)${escapeRegExp(marker)}`,
    'g'
  )
  const matches = text?.matchAll(pattern) || []
  const ids = Array.from(matches).map((v) => v[1])
  const tags = ids
    .map((id) => {
      const jump = `/${marker}${id}${marker}`
      return `${id}\t${path}\t${jump}`
    })
    .join('\n')
  return tags
}
