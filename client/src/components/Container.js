import React,{ useState, useEffect} from 'react'
import getProducts from '../services/productService';
import { useParams } from 'react-router-dom'


const Container = () => {
    const { category } = useParams()
    const [products, setProducts] = useState([]);

    useEffect(() => {
        const fetch = async () => {
            const res = await getProducts(category);
            console.log(res);
            setProducts(res);
        };
        fetch();
    }, []);

    const ProductCard = ({ pro }) => {
        return (
            <div className='Product'>
                <p> {pro.Name} </p>
                <p> {pro.Stock} </p>
            </div>
        );
    };

    const productRow = () => {
        if (products) {
            return products?.map((pro) => <ProductCard pro={pro} key={pro.Id}/>);
        }
    };

    return(
            <div className='Container'>{productRow()}</div>
    )
}

export default Container