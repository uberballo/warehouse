export const formatStockValue = (value) => {
  switch (value) {
    case 'OUTOFSTOCK':
      return 'Out of stock'
    case 'LESSTHAN10':
      return 'Less than 10'
    case 'INSTOCK':
      return 'Instock'
    default:
      return value
  }
}
