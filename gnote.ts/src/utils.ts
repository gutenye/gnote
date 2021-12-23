import fs from 'fs-jetpack'

export function emptyDir(dir: string): void {
  fs.list(dir)?.forEach((v) => fs.remove(`${dir}/${v}`))
}
