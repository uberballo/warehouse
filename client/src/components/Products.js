import React, { forwardRef } from 'react'
import AutoSizer from 'react-virtualized-auto-sizer'
import { FixedSizeGrid as Grid } from 'react-window'
import useWindowSize from '../hooks/windowSize'
import ProductCard from './ProductCard'

const Products = ({ products }) => {
  const windowSize = useWindowSize()
  const GUTTER_SIZE = 5
  const ROW_WIDTH = 250
  const ROW_HEIGHT = 100
  const COLUMN_COUNT = Math.floor(windowSize.width / ROW_WIDTH) || 0
  const ROW_COUNT = Math.ceil(products?.length / COLUMN_COUNT) || 0

  const Cell = ({ columnIndex, rowIndex, style }) => {
    const i = COLUMN_COUNT * rowIndex + columnIndex
    const pro = products[i]
    if (pro) {
      return (
        <div
          className={'GridItem'}
          style={{
            ...style,
            left: style.left + GUTTER_SIZE,
            top: style.top + GUTTER_SIZE,
            width: style.width - GUTTER_SIZE,
            height: style.height - GUTTER_SIZE,
          }}
        >
          <ProductCard pro={pro} />
        </div>
      )
    } else {
      return <div style={style}>Loading...</div>
    }
  }

  const innerElementType = forwardRef(({ style, ...rest }, ref) => (
    <div
      ref={ref}
      style={{
        ...style,
        paddingLeft: GUTTER_SIZE,
      }}
      {...rest}
    />
  ))
  return (
    <div className='Container' style={{ flex: '1 1 auto', height: '100vh' }}>
      <AutoSizer>
        {({ height, width }) => (
          <Grid
            columnCount={COLUMN_COUNT}
            columnWidth={ROW_WIDTH}
            height={height}
            innerElementType={innerElementType}
            rowCount={ROW_COUNT}
            rowHeight={ROW_HEIGHT}
            width={width}
          >
            {Cell}
          </Grid>
        )}
      </AutoSizer>
    </div>
  )
}

export default Products
