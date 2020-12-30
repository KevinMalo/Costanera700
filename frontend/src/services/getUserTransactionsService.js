const requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };

const url = "http://localhost"
const port ="9000"

function getUserTransactions(userId) {
    return fetch(`${url}:${port}/transaction/${userId}`, requestOptions)
    .then( res => res.json() )
    .then( res => res )
}

export default {
    getUserTransactions
}