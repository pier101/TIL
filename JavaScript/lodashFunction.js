import _ from 'lodash'

// 중복 제거
_.uniq([2,1,2]);
// uniqBy
// unionBy
// find
// findIndex
// remove
// cloneDeep


const usersA = [
    {userId: '1', name: 'kdw'},
    {userId: '2', name: 'iii'},
]
const usersB = [
    {userId: '1', name: 'kdw'},
    {userId: '3', name: 'ggg'},
]
const usersC = usersA.concat(usersB)
console.log('concat', usersC)
console.log('uniqBy', _.uniqBy(usersC, 'userId'))

const usersD = _unionBy(usersA, usersB, 'userId')
console.log('unionBy', usersD)

//--------------------

const users = [
    {userId: '1', name: 'kdw'},
    {userId: '2', name: 'iii'},
    {userId: '3', name: 'ggg'},
    {userId: '4', name: 'kkk'},
    {userId: '5', name: 'zzz'},
]

const foundUser = _.find(users, {name: 'ggg'})
const foundUserIndex = _.findIndex(users, {name: 'ggg'})
console.log(foundUser)
console.log(foundUserIndex)

_.remove(users, {name: 'zzz'})
console.log(users)