import {
  BrowserRouter as Router,
  Redirect,
  Route,
  Switch,
} from 'react-router-dom'
import ProductContainer from './components/ProductContainer'
import NavBar from './components/NavBar'

const App = () => {
  return (
    <div className='App'>
      <NavBar />
      <Router>
        <Switch>
          <Route path='/:category'>
            <ProductContainer />
          </Route>
        </Switch>
        <Route exact path='/'>
          <Redirect to='/facemasks' />
        </Route>
      </Router>
    </div>
  )
}

export default App
