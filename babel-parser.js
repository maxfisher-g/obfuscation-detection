#!/usr/bin/env node
/* eslint no-var: 0 */

const parser = require("@babel/parser");
const fs = require("fs");
const {parse} = require("@babel/parser");

const printTypes = false;

function walkJs(startNode, printTypes) {
    // walk the AST and print out any literals
    const n = startNode;

    if (printTypes) {
        console.log(`# type: ${n.type}`)
    }


    const loc = (n.loc != null) ? `[${n.loc.start.line},${n.loc.start.column}]` : "[]"
    switch (n.type) {
        case "File":
            walkJs(n.program)
            break
        case "ExpressionStatement":
            walkJs(n.expression)
            break
        case "VariableDeclaration":
            for (let i = 0; i < n.declarations.length; i++) {
                walkJs(n.declarations[i])
            }
            break
        case "VariableDeclarator":
            walkJs(n.id)
            walkJs(n.init)
            break
        case "Identifier":
            console.log(`{"type":"Identifier    ","value":"${n.name}","pos":${loc}}`)
            break
        case "StringLiteral":
            console.log(`{"type":"StringLiteral ","value":${JSON.stringify(n.value)},"pos":${loc},"extra":${JSON.stringify(n.extra)}}`)
            break
        case "NumericLiteral":
            console.log(`{"type":"NumericLiteral","value":${JSON.stringify(n.value)},"pos":${loc},"extra":${JSON.stringify(n.extra)}}`)
            break
        case "MemberExpression":
            walkJs(n.object)
            walkJs(n.property)
            break
        case "CallExpression":
            walkJs(n.callee)
            //console.log(n.callee)
            for (let i = 0; i < n.arguments.length; i++) {
                walkJs(n.arguments[i])
            }
            break
        case "AwaitExpression":
            walkJs(n.argument)
            break
        case "ArrayExpression":
            if (n.elements !== null) {
                for (let i = 0; i < n.elements.length; i++) {
                    walkJs(n.elements[i])
                }
            }
            break
        default:
            if (n.body !== undefined) {
                if (n.body.length === undefined) {
                    // FunctionDeclaration
                    walkJs(n.body)
                } else {
                    // Program, BlockStatement
                    for (let i = 0; i < n.body.length; i++) {
                        walkJs(n.body[i])
                    }
                }
            } else if (n.left !== undefined && n.right !== undefined) {
                // BinaryExpression, AssignmentExpression
                walkJs(n.left)
                walkJs(n.right)
            } else {
                console.log(`Found leaf node of type ${n.type} node @ ${loc}`);
                console.log(n)
            }

    }
}

function parseAndPrint(filename) {
    const file = fs.readFileSync(filename, "utf8");
    const ast = parser.parse(file);

    // walk the AST and print out any literals
    console.log(JSON.stringify(ast, null, "  "));

    walkJs(ast)
}

const filename = process.argv[2];
if (!filename) {
    console.error("no filename specified");
} else {
    parseAndPrint(filename)

}
