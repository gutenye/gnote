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
