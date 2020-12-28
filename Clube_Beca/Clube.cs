using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Clube_Beca
{
    public class Clube
    {
        private readonly ISocio _socio;

        public Clube(ISocio socio)
        {
            _socio = socio;
        }

        public double CalcularFrete(int socioID)
        {
            Socio socio = _socio.GetSocio(socioID);
            double valorFrete = socio.Limite * 0.5;

            return valorFrete;
        }
    }
}
