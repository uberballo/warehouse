import styled from 'styled-components'
import { Link } from 'react-router-dom'

const StyledNavBar = styled.div`
  display: flex;
  justify-content: space-around;

  position: fixed;
  top: 0px;
  z-index: 99;
  width: 100%;
  margin: 0 auto;

  padding: 10px 0;
  background-color: white;
`

const NavBar = () => {
  return (
    <StyledNavBar>
      <Link to='/facemasks'>face masks</Link>
      <Link to='/gloves'>gloves</Link>
      <Link to='/beanies'>beanies</Link>
    </StyledNavBar>
  )
}
export default NavBar
