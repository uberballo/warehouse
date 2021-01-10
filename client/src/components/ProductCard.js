import React from 'react'
import styled from 'styled-components'
import { formatStockValue } from '../util/stockValue'
import Row from './common/Row'

const StyledProductCard = styled.div`
  display: flex;
  flex-direction: column;
  padding: 5px;
  margin: 5px;
  border-style: groove;
`

const ProductCard = ({ pro }) => {
  return (
    <StyledProductCard>
      <p> {pro.Name} </p>
      <Row label='Manufacturer' value={pro.Manufacturer} />
      <Row label='Colors' value={pro.Color.join(', ')} />
      <Row label='Price' value={pro.Price} />
      <p> {formatStockValue(pro.Stock)} </p>
    </StyledProductCard>
  )
}

export default ProductCard
