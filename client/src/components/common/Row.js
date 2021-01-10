import React from 'react'
import styled from 'styled-components'

const StyledRow = styled.div`
  margin: 0px;
`
const Label = styled.p`
  float: left;
  margin: 2px;
`
const Value = styled.p`
  float: right;
  margin: 2px;
`

const Row = ({ label, value }) => {
  return (
    <StyledRow>
      <Label>{label}</Label>
      <Value>{value}</Value>
    </StyledRow>
  )
}

export default Row
