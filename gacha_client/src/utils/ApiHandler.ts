import axios from "axios";
import { cookies } from "next/headers";

const apiClient = axios.create({
    baseURL: process.env.NEXT_PUBLIC_API_URL,
})

async function getUserToken() {
    // get from local or cookie session
    const cookieStore = await cookies()
    const token = cookieStore.get('token')?.value

    if (!token) {
        throw new Error('Token not found, please login first')
    }

    return token
}

async function getUserTokenLogin(loginName: string, password: string) {
    const res = await apiClient.post('/login', {
        'loginName': loginName,
        'passwd': password,
    })
    return res.data.token
}

async function drawGacha(token: string, gachaCategId: string, amount: number) {
    const res = await apiClient.post('/draw', {
        'gachaId': gachaCategId,
        'amount': amount
    },
    {
        headers: {
            Authorization: `Bearer ${token}`,
            'Content-Type': 'application/json',
        },
    })
    return res.data
}

export {getUserToken, getUserTokenLogin, drawGacha};