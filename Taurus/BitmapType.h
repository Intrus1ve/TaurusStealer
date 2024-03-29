#pragma once
#include "Common.h"

using namespace Gdiplus;

class ImageC
{
public:
	void operator delete (void* in_pVoid);
	void* operator new(size_t in_size);
	void operator delete[](void* in_pVoid);
	void* operator new[](size_t in_size);

	friend class Brush;
	friend class TextureBrush;
	friend class Graphics;

	ImageC(IN const WCHAR* filename, IN BOOL useEmbeddedColorManagement = FALSE);
	ImageC(IN IStream* stream, IN BOOL useEmbeddedColorManagement = FALSE);
	virtual ~ImageC();

	static ImageC* FromFile(IN const WCHAR* filename, IN BOOL useEmbeddedColorManagement = FALSE);
	static ImageC* FromStream(IN IStream* stream, IN BOOL useEmbeddedColorManagement = FALSE);

	virtual ImageC* Clone();

	Status Save(IN const WCHAR* filename, IN const CLSID* clsidEncoder, IN const EncoderParameters *encoderParams = NULL);
	Status Save(IN IStream* stream, IN const CLSID* clsidEncoder, IN const EncoderParameters* encoderParams = NULL);

	Status SaveAdd(IN const EncoderParameters* encoderParams);
	Status SaveAdd(IN ImageC* newImage, IN const EncoderParameters* encoderParams);

	ImageType GetType() const;
	Status GetPhysicalDimension(OUT SizeF* size);
	Status GetBounds(OUT RectF* srcRect, OUT Unit* srcUnit);

	UINT GetWidth();
	UINT GetHeight();
	REAL GetHorizontalResolution();
	REAL GetVerticalResolution();
	UINT GetFlags();
	Status GetRawFormat(OUT GUID *format);
	PixelFormat GetPixelFormat();

	INT GetPaletteSize();
	Status GetPalette(OUT ColorPalette* palette, IN INT size);
	Status SetPalette(IN const ColorPalette* palette);

	ImageC* GetThumbnailImage(IN UINT thumbWidth, IN UINT thumbHeight, IN GetThumbnailImageAbort callback = NULL, IN VOID* callbackData = NULL);
	UINT GetFrameDimensionsCount();
	Status GetFrameDimensionsList(OUT GUID* dimensionIDs, IN UINT count);
	UINT GetFrameCount(IN const GUID* dimensionID);
	Status SelectActiveFrame(IN const GUID* dimensionID, IN UINT frameIndex);
	Status RotateFlip(IN RotateFlipType rotateFlipType);
	UINT GetPropertyCount();
	Status GetPropertyIdList(IN UINT numOfProperty, OUT PROPID* list);
	UINT GetPropertyItemSize(IN PROPID propId);
	Status GetPropertyItem(IN PROPID propId, IN UINT propSize, OUT PropertyItem* buffer);
	Status GetPropertySize(OUT UINT* totalBufferSize, OUT UINT* numProperties);
	Status GetAllPropertyItems(IN UINT totalBufferSize, IN UINT numProperties, OUT PropertyItem* allItems);
	Status RemovePropertyItem(IN PROPID propId);
	Status SetPropertyItem(IN const PropertyItem* item);

	UINT  GetEncoderParameterListSize(IN const CLSID* clsidEncoder);
	Status GetEncoderParameterList(IN const CLSID* clsidEncoder, IN UINT size, OUT EncoderParameters* buffer);

	Status GetLastStatus() const
	{
		Status lastStatus = lastResult;
		lastResult = Ok;

		return lastStatus;
	}

protected:
	ImageC() {}

	ImageC(GpImage *nativeImage, Status status)
	{
		SetNativeImage(nativeImage);
		lastResult = status;
	}

	VOID SetNativeImage(GpImage* nativeImage)
	{
		this->nativeImage = nativeImage;
	}

	_Post_equal_to_(status)
		Status SetStatus(Status status) const
	{
		if (status != Ok)
			return (lastResult = status);
		else
			return status;
	}

	GpImage* nativeImage;
	mutable Status lastResult;
	mutable Status loadStatus;

private:
	ImageC(IN const ImageC& C);
	ImageC& operator=(IN const ImageC& C);
};

class BitmapC : public ImageC
{
public:
	void operator delete(void* in_pVoid);
	void* operator new(size_t in_size);
	void operator delete[](void* in_pVoid);
	void* operator new[](size_t in_size);

	friend class ImageC;
	friend class CachedBitmap;

	BitmapC(IN const WCHAR *filename, IN BOOL useEmbeddedColorManagement = FALSE);
	BitmapC(IN IStream *stream, IN BOOL useEmbeddedColorManagement = FALSE);

	static BitmapC* FromFile(IN const WCHAR *filename, IN BOOL useEmbeddedColorManagement = FALSE);
	static BitmapC* FromStream(IN IStream *stream, IN BOOL useEmbeddedColorManagement = FALSE);

	BitmapC(_In_ INT width, _In_ INT height, _In_ INT stride, _In_ PixelFormat format, BYTE* scan0);
	BitmapC(IN INT width,IN INT height, IN PixelFormat format = PixelFormat32bppARGB);
	BitmapC(IN INT width,IN INT height, IN  Graphics* target);

	BitmapC* Clone(IN const Rect& rect, IN PixelFormat format);
	BitmapC* Clone(IN INT x, IN INT y, IN INT width, IN INT height, IN PixelFormat format);
	BitmapC* Clone(IN const RectF& rect, IN PixelFormat format);
	BitmapC* Clone(IN REAL x, IN REAL y, IN REAL width, IN REAL height, IN PixelFormat format);

	Status LockBits(IN const Rect* rect, IN UINT flags, IN PixelFormat format, OUT BitmapData* lockedBitmapData);
	Status UnlockBits(IN BitmapData* lockedBitmapData);
	Status GetPixel(IN INT x, IN INT y, OUT Color *color);
	Status SetPixel(IN INT x, IN INT y, IN const Color &color);
	Status SetResolution(IN REAL xdpi, IN REAL ydpi);

	BitmapC(IN IDirectDrawSurface7* surface);
	BitmapC(IN const BITMAPINFO* gdiBitmapInfo, IN VOID* gdiBitmapData);
	BitmapC(IN HBITMAP hbm, IN HPALETTE hpal);

	BitmapC(IN HICON hicon);
	BitmapC(IN HINSTANCE hInstance, IN const WCHAR * bitmapName);
	static BitmapC* FromDirectDrawSurface7(IN IDirectDrawSurface7* surface);
	static BitmapC* FromBITMAPINFO(IN const BITMAPINFO* gdiBitmapInfo, IN VOID* gdiBitmapData);
	static BitmapC* FromHBITMAP(IN HBITMAP hbm, IN HPALETTE hpal);
	static BitmapC* FromHICON(IN HICON hicon);
	static BitmapC* FromResource(IN HINSTANCE hInstance, IN const WCHAR * bitmapName);

	Status GetHBITMAP(IN const Color& colorBackground, OUT HBITMAP *hbmReturn);
	Status GetHICON(HICON *hicon);
private:
	BitmapC(const BitmapC &);
	BitmapC& operator=(const BitmapC &);

protected:
	BitmapC(GpBitmap *nativeBitmap);
};