#!/usr/bin/env bun
const dir = __dirname;

import path from 'node:path';
const root = path.resolve(dir, '../../../');
const env = path.resolve(root, '.env');

require('dotenv').config({ path: env });
const program = process.env.B as string;

const target = path.resolve(dir, '..', 'bin', program);
const dest = path.resolve(root, 'bin', program);

import fs from 'node:fs';
if(!fs.existsSync(target))
    throw new Error(`FATAL The target file '${ target }' doesn't exist`);
else if(!fs.existsSync(dest))
    await Bun.$`mv ${ target } ${ dest }`;
else
    await Bun.$`rm ${ dest } && cp ${ target } ${ dest }`;