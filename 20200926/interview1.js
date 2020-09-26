// function exchange(a, b) {
//     var c = a;
//     a = b;
//     b = c;
//     a = a + 'world'; 
//     b = b + 'test';
//     console.log(a);
//     console.log(b);
// }

// var a = 'hello,'; 
// var b = 'test';

// exchange(a, b);

// console.log(a);
// console.log(b);


var a = { name: 'hello' }, b = { name: 'hi' }
function exchange(a, b) {
    var c = b; b = a; a = c;
    a.name = a.name + "1";
    b.name = b.name + "2";
    console.info(a.name, b.name);
}
exchange(a, b);
console.info(a.name, b.name);
