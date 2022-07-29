// one dimention arrays

const array = [10, 20, 23, 1, 31, 100, 21, 101, 64, 12];

/**
 * @param {array<number>} arr
 */
const reduceArr = (arr) => arr.reduce((prev, currentVal) => prev + currentVal);
// console.log(reduceArr(array));

/**
 * @description get an array and a 'n', returns the first n items in that array
 * @param {array<number>} arr
 * @param {number} n - number of items to take of the array, in sequential order, starting from index zero[0]
 */
const firstN = (arr, n = 0) => {
  if (arr.length === 0) return [];
  if (n <= 1) return arr[0];

  const newArr = [];
  for (let i = 0; i < n; i++) {
    if (!arr[i]) break;

    newArr.push(arr[i]);
  }
  return newArr;
};

console.log(firstN(array, 20));
