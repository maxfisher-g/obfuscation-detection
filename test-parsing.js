function test() {
    var mystring1 = "hello1";
    var mystring2 = 'hello2';
    var mystring3 = "hello'3'";
    var mystring4 = 'hello"4"';
    var mystring5 = "hello\"5\"";
    var mystring6 = "hello\'6\'";
    var mystring7 = 'hello\'7\'';
    var mystring8 = "hello" + "8";
    var mystring9 = `hello9`;
    var mystring10 = `hello"'${10}"'`;
    var mystring11 = `hello
//"'11"'`;
    var mystring12 = `hello"'${5.6 + 6.4}"'`;
}

function test2(param1, param2, param3 = "ahd") {
    return param1 + param2 + param3
}

function test3(a, b, c) {
    for (var i = a; i < b; i++) {
outer:
        for (var j = 1; j < 3; j++) {
            for (var k = j; k < j + 10; k++) {
                if (j === 2) {
                    break outer
                }
            }
        }
        c = c * i
        if (c % 32 === 0) {
            continue
        }
        console.log("here")
    }
    console.log("End")
}

function test4() {
    const a = [1, 2, 3]
    try {
        if (a[1] === 3) {
            console.log(a[-1])
        } else if (a[1] === 2) {
            console.log(a[1])
        } else {
            console.log(a[2])
        }
    } catch (e) {
        var f = "abc"
        console.log(e + f)
    }

    switch (a[0]) {
        case 1:
            console.log("Hp")
            break
        default:
            console.log("Hq")
            break
    }
}

// unnamed
let Rectangle = class {
    constructor(height, width) {
        this.height = height;
        this.width = width;
    }
};
console.log(Rectangle.name);
// output: "Rectangle"

// named
Rectangle = class Rectangle2 {
    constructor(height, width) {
        this.height = height;
        this.width = width;
    }
};
console.log(Rectangle.name);
// output: "Rectangle2"