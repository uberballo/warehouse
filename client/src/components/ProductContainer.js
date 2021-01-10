import React from 'react'
import { useParams } from 'react-router-dom'
import styled from 'styled-components'
import useSWR from 'swr'
import * as constants from '../constants'
import getProducts from '../services/productService'
import Products from './Products'

const StyledProductContainer = styled.div`
  text-align: center;
  justify-content: center;
`

const ProductContainer = () => {
  const { category } = useParams()
  const { data, error } = useSWR(category, getProducts, {
    refreshInterval: constants.REFRESH_INTERVAL_IN_MS,
  })

  if (error) return <div>failed to load</div>
  if (!data) return <div>loading...</div>

  return (
    <StyledProductContainer>
      <Products products={data} />
    </StyledProductContainer>
  )
}

export default ProductContainer
