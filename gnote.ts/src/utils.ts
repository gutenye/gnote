import fsSync from 'fs'
import fs from 'fs/promises'
import path from 'path'
import { fileURLToPath } from 'url'

export function exit(message: string): void {
  console.log(message)
  process.exit(1)
}

export function readPkg(): { [key: string]: any } {
  return JSON.parse(
    fsSync.readFileSync(
      path.resolve(sourceFileDirname(), '../package.json'),
      'utf8'
    )
  )
}

export function sourceFileDirname(): string {
  return path.dirname(fileURLToPath(import.meta.url))
}

export async function writeFileWithMkdir(
  file: string,
  data: string
): Promise<void> {
  await fs.mkdir(path.dirname(file), { recursive: true })
  await fs.writeFile(file, data)
}
