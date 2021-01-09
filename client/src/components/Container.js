import React, { useState, useEffect, forwardRef } from "react";
import getProducts from "../services/productService";
import { useParams } from "react-router-dom";
import { FixedSizeGrid as Grid } from "react-window";
import AutoSizer from "react-virtualized-auto-sizer";
import useWindowSize from '../hooks/windowSize'


const Container = () => {
  const { category } = useParams();
  const [products, setProducts] = useState([]);
  const windowSize = useWindowSize()
  const GUTTER_SIZE = 5;
  const ROW_WIDTH = 250;
  const ROW_HEIGHT = 100;
  const COLUMN_COUNT = Math.floor((windowSize.width / ROW_WIDTH));
    console.log("count",COLUMN_COUNT)

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
      <div className="Product">
        <p> {pro.Name} </p>
        <p> {pro.Stock} </p>
      </div>
    );
  };


  const Cell = ({ columnIndex, rowIndex, style }) => {
    const i = COLUMN_COUNT* rowIndex + columnIndex;
    const pro = products[i];
    if (pro) {
      return (
        <div
          className={"GridItem"}
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
      );
    } else {
      return <div style={style}>Loading...</div>;
    }
  };

  const innerElementType = forwardRef(({ style, ...rest }, ref) => (
    <div
      ref={ref}
      style={{
        ...style,
        paddingLeft: GUTTER_SIZE,
      }}
      {...rest}
    />
  ));

  return (
    <div className="Container" style={{ flex: "1 1 auto", height: "100vh" }}>
      <AutoSizer>
        {({ height, width }) => (
          <Grid
            columnCount={COLUMN_COUNT}
            columnWidth={ROW_WIDTH}
            height={height}
            innerElementType={innerElementType}
            rowCount={200}
            rowHeight={ROW_HEIGHT}
            width={width}
          >
            {Cell}
          </Grid>
        )}
      </AutoSizer>
    </div>
  );
};

export default Container;
