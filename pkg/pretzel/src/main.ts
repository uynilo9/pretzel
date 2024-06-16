const args = require('yargs')(process.argv.slice(2)).parse();
if(args.allowed === "false" || args.allowed === undefined)
    process.exit(1)

console.log(args)