import { Container, Stack } from '@chakra-ui/react'
import Navbar from './components/Navbar'
import TodoForm from './components/TodoForm'

// make todoitems and todolists 
function App() {
  return (
    <Stack h="100vh">
      <Navbar />
      <Container>
        <TodoForm>

        </TodoForm>
      </Container>
    </Stack>
  )
}

export default App;
