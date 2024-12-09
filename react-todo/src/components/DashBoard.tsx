import { useState, FormEvent } from 'react'
import { CheckBadgeIcon, ArrowPathIcon, ShieldCheckIcon } from '@heroicons/react/24/solid'
import { useMutateAuth } from '../hooks/useMutateAuth'

export const DashBoard = () => {
  return (
    <div className="flex justify-center items-center flex-col min-h-screen text-gray-600 font-mono">
      <div className="flex items-center my-3">
        <ShieldCheckIcon className="h-8 w-8 mr-3 text-indigo-500 cursor-pointer" />
        <span className="text-center text-3xl font-extrabold">
          ダッシュボード
        </span>
      </div>
      <div className="flex items-center my-3">
        <a href='/todo'>
          <button className="text-xl bg-blue-500 hover:bg-blue-700 text-white font-bold max-w-full w-72 py-2 px-10 rounded">
            ユーザ検索
          </button>
        </a>
      </div>
      <div className="flex items-center my-3">
        <a href='/todo'>
          <button className="text-xl bg-blue-500 hover:bg-blue-700 text-white font-bold max-w-full w-72 py-2 px-10 rounded">
            Todoリスト
          </button>
        </a>
      </div>

    </div>
  )
}