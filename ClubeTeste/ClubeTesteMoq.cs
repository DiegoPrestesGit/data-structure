using Clube_Beca;
using Microsoft.VisualStudio.TestTools.UnitTesting;
using Moq;

namespace ClubeTeste
{
    [TestClass]
    public class ClubeTesteMoq
    {
        private readonly Clube _target;
        private readonly Mock<ISocio> _mock;
    }
}
