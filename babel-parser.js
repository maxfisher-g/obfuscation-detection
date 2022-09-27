#!/usr/bin/env node
/* eslint no-var: 0 */

const parser = require("@babel/parser");
const fs = require("fs");

const printTypes = false;

function locationString(node) {
    return (node.loc != null) ? `[${node.loc.start.line},${node.loc.start.column}]` : "[]"
}

const parseOutputLines = []

function logJSON(type, subtype, name, pos, array, extra = null) {
    const extraValue = (extra !== null) ? `${JSON.stringify(extra)}` : "{}"
    const arrayValue = (array !== null) ? array : false
    const json = `{"type":"${type}","subtype":"${subtype}","data":${JSON.stringify(name)},"pos":${pos},` +
        `"array":${arrayValue}, "extra":${extraValue}}`
    parseOutputLines.push(json)
}
function logIdentifierJSON(subtype, name, pos, extra = null) {
    logJSON("Identifier", subtype, name, pos, null, extra)
}

function logLiteralJSON(subtype, value, pos, inArray, extra = null) {
    logJSON("Literal", subtype, value, pos, inArray, extra)
}

function multiWalkAst(nodeList, isInArray = false) {
    if (nodeList !== null) {
        for (let i = 0; i < nodeList.length; i++) {
            walkAst(nodeList[i], isInArray)
        }
    }
}

function walkAst(startNode, isInArray = false) {
    // walk the AST and print out any literals
    const n = startNode;

    if (printTypes) {
        console.log(`# type: ${n.type}`)
    }

    const loc = locationString(n)
    switch (n.type) {
        case "File":
            walkAst(n.program)
            break
        case "Program":
        // fall-through
        case "BlockStatement":
            multiWalkAst(n.body, isInArray)
            break
        case "ArrowFunctionExpression":
        // fall-through
        case "FunctionDeclaration":
            if (n.id !== null) {
                logIdentifierJSON("Function", n.id.name, locationString(n.id))
            }
            walkAst(n.body)
            break
        case "AwaitExpression":
            walkAst(n.argument)
            break
        case "AssignmentExpression":
        // fall-through
        case "BinaryExpression":
            walkAst(n.left)
            walkAst(n.right)
            break
        case "ExpressionStatement":
            walkAst(n.expression)
            break
        case "CallExpression":
            walkAst(n.callee)
            multiWalkAst(n.arguments, isInArray)
            break
        case "MemberExpression":
            walkAst(n.object)
            walkAst(n.property)
            break
        case "VariableDeclaration":
            multiWalkAst(n.declarations, isInArray)
            break
        case "VariableDeclarator":
            logIdentifierJSON("Variable", n.id.name, locationString(n.id))
            walkAst(n.init)
            break
        case "Identifier":
            logJSON("OtherIdentifier", n.name, loc)
            break
        case "StringLiteral":
            logLiteralJSON("String", n.value, loc, isInArray, n.extra)
            break
        case "NumericLiteral":
            logLiteralJSON("Numeric", n.value, loc, isInArray, n.extra)
            break
        case "ArrayExpression":
            multiWalkAst(n.elements, true)
            break
        case "TemplateLiteral":
            multiWalkAst(n.quasis, isInArray)
            multiWalkAst(n.expressions, isInArray)
            break
        case "TemplateElement":
            logLiteralJSON("StringTemplate", n.value.raw, loc, isInArray, n.value)
            break
        default:
            console.log(`Found unknown node of type ${n.type} node @ ${loc}`);
            console.log(n)
    }
}

function parseAndPrint(filename) {
    const file = fs.readFileSync(filename, "utf8");
    const ast = parser.parse(file);

    // walk the AST and print out any literals
    //console.log(JSON.stringify(ast, null, "  "));

    walkAst(ast)
    allJson = "[\n" + parseOutputLines.join(",\n") + "\n]"
    console.log(allJson)
}

const filename = process.argv[2];
if (!filename) {
    console.error("no filename specified");
} else {
    parseAndPrint(filename)

}
