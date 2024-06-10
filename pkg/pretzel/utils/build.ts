#!/usr/bin/env bun
const dir = __dirname;

import os from 'node:os';
const arch = os.arch();
const platform = os.platform();

import path from 'node:path';
const root = path.resolve(dir, '../../../');
const env = path.resolve(root, '.env');

require('dotenv').config({ path: env });
const program = process.env.B as string;


const source = path.resolve(dir, '..', 'src', 'main.ts');
const dest = path.resolve(dir, '..', 'bin', program);

let target: string;

if(platform === 'darwin')
    target = arch === 'arm64' ? 'bun-darwin-arm64' : 'bun-darwin-x64';
else if(platform === 'linux')
    target = arch === 'arm64' ? 'bun-linux-arm64' : 'bun-linux-x64';
else if(platform === 'win32')
    target = 'bun-windows-x64';
else
    throw new Error('WARN Not implemented yet');

await Bun.$`bun build --compile ${ source } --target=${ target } --outfile=${ dest }`;