function App() {
  return (
    <div className="flex items-center justify-center bg-[#050505] w-full h-screen font-mono selection:bg-white selection:text-black">
      <section className="w-full max-w-2xl px-6">
        <div className="mb-12 text-center">
          <h1 className="text-4xl font-bold text-white tracking-widest uppercase">
            [Skimdex]
          </h1>
        </div>

        <form action="/search" method="GET" className="relative group">
          <div className="absolute -inset-0.5 bg-zinc-800 rounded-sm opacity-50 group-focus-within:bg-white transition-all duration-300"></div>
          <div className="relative flex items-center bg-[#0a0a0a] rounded-sm overflow-hidden">
            <span className="pl-4 text-zinc-600 text-lg select-none">&gt;</span>
            <input
              type="text"
              name="q"
              placeholder="ENTER_QUERY..."
              className="w-full bg-transparent text-white text-lg p-2 focus:outline-none placeholder-zinc-800"
              autoFocus
              autoComplete="off"
            />
            <button
              type="submit"
              className="px-2 py-3 bg-zinc-900 text-zinc-500 hover:bg-white hover:text-black transition-all duration-200 font-bold border-l border-zinc-800"
            >
              🔍
            </button>
          </div>
        </form>
      </section>
    </div>
  );
}

export default App;
