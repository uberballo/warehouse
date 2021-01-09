import React,{ useState, useEffect, forwardRef} from 'react'
import getProducts from '../services/productService';
import { useParams } from 'react-router-dom'
import { FixedSizeGrid as Grid } from 'react-window';
import AutoSizer from "react-virtualized-auto-sizer";





const Container = () => {
    const { category } = useParams()
    const [products, setProducts] = useState([]);

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
            <div className='Product'>
                <p> {pro.Name} </p>
                <p> {pro.Stock} </p>
            </div>
        );
    };

    const productRow = () => {
        if (products) {
            return products?.map((pro) => <ProductCard pro={pro} key={pro.Id}/>);
        }
    };

function rowRenderer({key, index, style}) {
    return (
      <div key={key} >
          <ProductCard pro={products[index]}/>
      </div>
    );
  }
  const GUTTER_SIZE = 50;

  const rows = 5
  const tCell = (props) => {
      console.log(props)
        return <div >
            Loading...
        </div>
  }
  const Cell = ({ columnIndex, rowIndex, style}) => {
      const i = rows*rowIndex + columnIndex
      //<ProductCard pro={products[i]}/>
      const pro = products[i]
      if (pro){
      return(
    <div 
    className={"GridItem"}

    style={{
        ...style,
        left: style.left + GUTTER_SIZE,
        top: style.top + GUTTER_SIZE,
        width: style.width - GUTTER_SIZE,
        height: style.height - GUTTER_SIZE,
      }
    }>
                <p> {pro.Name} </p>
                <p> {pro.Manufacturer} </p>
                <p> {pro.Stock} </p>
    </div>)}else{
        return <div style={style}>
            Loading...
        </div>
    }
  }
   
/*
 return(
            <div className='Container'>{productRow()}</div>
    )
    */
   const innerElementType = forwardRef(({ style, ...rest }, ref) => (
    <div
      ref={ref}
      style={{
        ...style,
        paddingLeft: GUTTER_SIZE,
        paddingTop: GUTTER_SIZE
      }}
      {...rest}
    />
  ));
  
 return(

            <div className='Container' style={{ flex: '1 1 auto' , height: '100vh'}} >
                 <AutoSizer>
    {({ height, width }) => (
    <Grid
    columnCount={width/250}
    columnWidth={250}
    height={height}
        innerElementType={innerElementType}

    rowCount={100}
    rowHeight={100}
    width={width}
  >
    {Cell}
  </Grid>
    )}
    </AutoSizer>
  </div>
    )
}

export default Container