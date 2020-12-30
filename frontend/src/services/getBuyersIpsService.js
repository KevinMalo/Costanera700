const requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };

const url = "http://localhost"
const port ="9000"

function getBuyersIps(userId) {
    return fetch(`${url}:${port}/buyers-ip/${userId}`, requestOptions)
    .then( res => res.json() )
    .then( res => res )
}

export default {
    getBuyersIps
}