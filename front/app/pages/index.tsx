import Link from 'next/link'
import Counter from '../components/counter'

export const Home = (): JSX.Element => (
  <div>
  <Link href='/posts/hoge'>
  <a>hoge</a>
  </Link>
  <Counter></Counter>
  </div>
)

export default Home
