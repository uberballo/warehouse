import React, { forwardRef } from 'react'
import AutoSizer from 'react-virtualized-auto-sizer'
import { FixedSizeGrid as Grid } from 'react-window'
import * as constants from '../constants'
import useWindowSize from '../hooks/windowSize'
import ProductCard from './ProductCard'
import styled from 'styled-components'

const Products = ({ products }) => {
  const windowSize = useWindowSize()
  const columnCount = Math.floor(windowSize.width / constants.ROW_WIDTH) || 0
  const rowCount = Math.ceil(products?.length / columnCount) || 0
  const gutter_size = constants.GUTTER_SIZE

  const Cell = ({ columnIndex, rowIndex, style }) => {
    const index = columnCount * rowIndex + columnIndex
    const product = products[index]

    if (product) {
      return (
        <div
          className={'GridItem'}
          style={{
            ...style,
            left: style.left + gutter_size,
            top: style.top + gutter_size,
          }}
        >
          <ProductCard pro={product} />
        </div>
      )
    } else {
      return <div style={style}>Loading...</div>
    }
  }

  //This height is required for the autosizer.
  const StyledProducts = styled.div`
    height: 97vh;
  `
  const style = {}
  return (
    <StyledProducts>
      <AutoSizer style={style}>
        {({ height, width }) => (
          <Grid
            columnCount={columnCount}
            columnWidth={constants.ROW_WIDTH}
            height={height}
            rowCount={rowCount}
            rowHeight={constants.ROW_HEIGHT}
            width={width}
            style={style}
          >
            {Cell}
          </Grid>
        )}
      </AutoSizer>
    </StyledProducts>
  )
}

export default Products
