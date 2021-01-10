import React from 'react'
import styled from 'styled-components'

const StyledProductCard = styled.div`
  padding: 5px;
  margin: 5px;
  border-style: groove;
`

const ProductCard = ({ pro }) => {
  return (
    <StyledProductCard>
      <p> {pro.Name} </p>
      <p> {pro.Manufacturer} </p>
      <p> {pro.Color}</p>
      <p> {pro.Price}</p>
      <p> {pro.Stock} </p>
    </StyledProductCard>
  )
}

export default ProductCard
