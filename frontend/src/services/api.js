const requestOptions = {
    method: 'GET',
    redirect: 'follow'
  };

const url = "http://localhost"
const port ="9000"

function getBuyers() {
    return fetch(`${url}:${port}/buyers`, requestOptions)
    .then( res => res.json() )
    .then( res => res.buyers )
}

export default {
    getBuyers
}