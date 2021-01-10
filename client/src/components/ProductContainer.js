import React from 'react'
import { useParams } from 'react-router-dom'
import styled from 'styled-components'
import useSWR from 'swr'
import * as constants from '../constants'
import getProducts from '../services/productService'
import Products from './Products'

const StyledProductContainer = styled.div`
  text-align: center;
`

const StyledLoadingErrorContainer = styled.div`
  text-align: center;
  padding: 70px 0;
`

const ProductContainer = () => {
  const { category } = useParams()
  const { data, error } = useSWR(category, getProducts, {
    refreshInterval: constants.REFRESH_INTERVAL_IN_MS,
  })

  if (error)
    return (
      <StyledLoadingErrorContainer>
        Failed to load products
      </StyledLoadingErrorContainer>
    )
  if (!data)
    return <StyledLoadingErrorContainer>Loading</StyledLoadingErrorContainer>

  return (
    <StyledProductContainer>
      <Products products={data} />
    </StyledProductContainer>
  )
}

export default ProductContainer
