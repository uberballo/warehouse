import axios from 'axios';

const baseUrl = '/api/products/';

const getProducts = async (category) => {
    const res = await axios.get(`${baseUrl}${category}`);
    console.log(res);
    return res.data.data;
};

export default getProducts;
