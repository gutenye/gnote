import fs from 'fs/promises'
import path from 'path'

export function exit(message: string): void {
  console.log(message)
  process.exit(1)
}

export async function writeFileWithMkdir(
  file: string,
  data: string
): Promise<void> {
  await fs.mkdir(path.dirname(file), { recursive: true })
  await fs.writeFile(file, data)
}

export async function emptyDir(dir: string): Promise<void> {
  const files = await fs.readdir(dir)
  for (const file of files) {
    await fs.rm(`${dir}/${file}`, { recursive: true })
  }
}
