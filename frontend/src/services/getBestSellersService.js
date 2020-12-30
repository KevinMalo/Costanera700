const requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };

const url = "http://localhost"
const port ="9000"

function getBestSellers() {
    return fetch(`${url}:${port}/best-sellers`, requestOptions)
    .then( res => res.json() )
    .then( res => res)
}

export default {
    getBestSellers
}