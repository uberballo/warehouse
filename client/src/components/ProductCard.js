import React from 'react'
  const ProductCard = ({ pro }) => {
    return (
      <div className="Product">
        <p> {pro.Name} </p>
        <p> {pro.Stock} </p>
      </div>
    );
  };

  export default ProductCard