import './App.css';
import getProducts from './services/productService';
import { useEffect, useState } from 'react';
import useSWR from 'swr';

const App = () => {
    const [products, setProducts] = useState([]);

    useEffect(() => {
        const fetch = async () => {
            const res = await getProducts('gloves');
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
            return products?.map((pro) => <ProductCard pro={pro} />);
        }
    };
    return (
        <div className='App'>
            <div className='Navbar'>
                <a href='/facemasks'>masks</a>
            </div>
            <div className='Container'>{productRow()}</div>
        </div>
    );
};

export default App;
