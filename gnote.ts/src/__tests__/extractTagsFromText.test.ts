import extractTagsFromText from '../extractTagsFromText'

test('works', () => {
  const result = extractTagsFromText({
    text: `
hello world
foo *bar*
link step
*baz* and 
		`,
    path: '/note',
    marker: '*',
  })
  expect(result).toEqual(
    `
bar\t/note\t/*bar*
baz\t/note\t/*baz*
	`.trim()
  )
})
