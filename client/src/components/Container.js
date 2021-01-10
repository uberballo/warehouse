import React from 'react'
import { useParams } from 'react-router-dom'
import useSWR from 'swr'
import getProducts from '../services/productService'
import Products from './Products'

const Container = () => {
  const { category } = useParams()
  const { data, error } = useSWR(category, getProducts, {
    refreshInterval: 5000,
  })

  if (error) return <div>failed to load</div>
  if (!data) return <div>loading...</div>

  return <Products products={data} />
}

export default Container
