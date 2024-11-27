import { file } from 'bun';
import { join, resolve } from 'path';

/**
 * Utility to help parse files for AoC. Will default to `input.txt` if no
 * argument is provided
 *
 * @param context the dir from which to resolve input path
 * @returns Promise<string> that will resolve to the file contents
 */
export async function readFileAsText(context: string) {
  const filePath = process.argv.length < 3 ? 'input.txt' : process.argv.at(-1);
  if (!filePath) {
    throw new Error('Failed to find input file');
  }
  const path = resolve(join(context, filePath));
  return (await file(path).text()).trim();
}
