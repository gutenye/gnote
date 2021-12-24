import alias from 'alias-hq'

export default {
  transform: { '^.+\\.tsx?$': ['esbuild-jest', { format: 'esm' }] },
  extensionsToTreatAsEsm: ['.ts'],
  moduleNameMapper: alias.get('jest'),
}
