import styled from 'styled-components'

const StyledNavBar = styled.div`
  text-align: center;
  position: fixed;
  top: 0px;
  width: 100%;
`

const NavBar = () => {
  return (
    <StyledNavBar>
      <a href='/facemasks'>face masks</a>
      <a href='/gloves'>gloves</a>
      <a href='/beanies'>beanies</a>
    </StyledNavBar>
  )
}
export default NavBar
