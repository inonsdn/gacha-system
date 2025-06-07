import {getUserTokenLogin, drawGacha} from '@/utils/ApiHandler'

export const dynamic = 'force-dynamic' // Always run on server

export default async function GachaPage() {

    const token = await getUserTokenLogin('nonser', 'testpass')

    let data = '{}'
    if (!token) {
        data = '{}'
    } else {
        data = await drawGacha(token, '67c61611-0164-4f92-96ec-0b3ffc0c3b3b', 10)
    }


    return (
    <main>
        <h2>Draw Result</h2>
        <pre>{JSON.stringify(data, null, 2)}</pre>
    </main>
    )
}
