import { useEffect, useState } from 'react'

interface Curso {
  id: string;
  nombre: string;
}

function App() {
  const [cursos, setCursos] = useState<Curso[]>([]);
  const [loading, setLoading] = useState<boolean>(true); 
  const [error, setError] = useState<string | null>(null); 

  useEffect(() => {

    fetch('/api/curso').then(async (response) => {
        if (!response.ok) {
          throw new Error('Error al conectar con el servidor Go');
        }
        return response.json();
      })
      .then((data) => {
        setCursos(data); 
        setLoading(false); 
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false); 
      });
  }, []);

  if (loading) {
    return <div style={{ padding: '20px' }}> Cargando cursos desde el backend</div>;
  }

  if (error) {
    return <div style={{ padding: '20px', color: 'red' }}> Error: {error}</div>;
  }

  if (cursos.length === 0) {
    return <div style={{ padding: '20px' }}>No hay cursos disponibles.</div>;
  }

  return (
    <div style={{ padding: '20px', fontFamily: 'sans-serif' }}>
      <h1>Lista de Cursos</h1>
      <ul>
        {cursos.map((curso) => (
          <li key={curso.id} style={{ margin: '10px 0' }}>
            <strong>ID {curso.id}:</strong> {curso.nombre}
          </li>
        ))}
      </ul>
    </div>
  )
}

export default App