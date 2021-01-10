import styled from 'styled-components'

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
      <a href='/facemasks'>face masks</a>
      <a href='/gloves'>gloves</a>
      <a href='/beanies'>beanies</a>
    </StyledNavBar>
  )
}
export default NavBar
