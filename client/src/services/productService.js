import axios from 'axios'

const baseUrl = '/api/products/'

const getProducts = async (category) => {
  const res = await axios.get(`${baseUrl}${category}`)
  console.log(res.data)
  if (res.status === 200) {
    return res.data
  }
}

export default getProducts
